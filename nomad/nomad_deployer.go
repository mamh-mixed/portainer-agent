package nomad

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/hashicorp/nomad/api"
	nomadapi "github.com/hashicorp/nomad/api"
	"github.com/pkg/errors"
)

// Deployer represents a service to deploy resources inside a Nomad environment.
type Deployer struct {
	client *nomadapi.Client
}

// NewDeployer initializes a new Nomad api client.
func NewDeployer() (*Deployer, error) {
	//DefaultConfig will try to retrieve NOMAD_ADDR and NOMAD_TOKEN from ENV
	client, err := nomadapi.NewClient(nomadapi.DefaultConfig())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init nomad api client")
	}
	return &Deployer{client: client}, nil
}

// Deploy attempts to run a Nomad job via provided job file
func (d *Deployer) Deploy(ctx context.Context, name string, filePaths []string, prune bool) error {
	if len(filePaths) == 0 {
		return errors.New("missing Nomad job file paths")
	}
	jobFile, err := ioutil.ReadFile(filePaths[0])
	if err != nil {
		return errors.Wrap(err, "failed to read Nomad job file")
	}
	job, err := d.client.Jobs().ParseHCL(string(jobFile), true)

	// Force the region to be that of the job.
	if r := job.Region; r != nil {
		d.client.SetRegion(*r)
	}

	// Force the namespace to be that of the job.
	if n := job.Namespace; n != nil {
		d.client.SetNamespace(*n)
	}

	// Check if the job is periodic or is a parameterized job
	periodic := job.IsPeriodic()
	paramjob := job.IsParameterized()
	multiregion := job.IsMultiregion()

	// Set the register options
	runOpts := &api.RegisterOptions{
		PolicyOverride: false,
		PreserveCounts: false,
		EvalPriority:   0,
	}

	// Submit the job
	_, _, err = d.client.Jobs().RegisterOpts(job, runOpts, nil)
	if err != nil {
		return errors.Wrap(err, "failed to run nomad job")
	}

	if periodic || paramjob || multiregion {
		if periodic && !paramjob {
			loc, err := job.Periodic.GetLocation()
			if err == nil {
				now := time.Now().In(loc)
				if _, err := job.Periodic.Next(now); err != nil {
					return errors.Wrap(err, "failed to run nomad periodic job")
				}
			}
		}
	}

	return nil
}

// Remove attempts to stop a Nomad job via provided job name/id
func (d *Deployer) Remove(ctx context.Context, name string, filePaths []string) error {
	jobID := strings.TrimSpace(name)
	jobs, _, err := d.client.Jobs().PrefixList(jobID)
	if err != nil {
		return errors.Wrap(err, "failed to list jobs with prefix")
	}
	if len(jobs) == 0 {
		return errors.New(fmt.Sprintf("no job(s) with prefix or id %s found", jobID))
	}
	q := &api.QueryOptions{Namespace: jobs[0].JobSummary.Namespace}
	job, _, err := d.client.Jobs().Info(jobs[0].ID, q)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve nomad job info")
	}

	opts := &api.DeregisterOptions{Purge: true}
	wq := &api.WriteOptions{Namespace: jobs[0].JobSummary.Namespace}
	_, _, err = d.client.Jobs().DeregisterOpts(*job.ID, opts, wq)
	if err != nil {
		return errors.Wrap(err, "failed to purge nomad job")
	}
	return nil
}
