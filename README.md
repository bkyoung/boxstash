# Boxstash
Boxstash provides an implementation of the [Vagrant Cloud API](https://www.vagrantup.com/docs/vagrant-cloud/api.html) for privately hosting vagrant boxes.  It serves the specific use case of allowing vagrant end-users to download box images managed by the API service, which can be hosted publicly or privately as the service administrator sees fit.  However, unlike Vagrant Cloud, this service is not designed to be a multi-tenant, self-service API.  Nearly all of the functionality the API provides is only accessible to the admin user.  That is to say, the administrator is who may create namespaces (organizations/users in boxstash parlance), new boxes, new box versions, new box providers, release, revoke, upload, and delete.

## Local Environment
This project is set up for efficient local development that gives the developer a decent sense of how the software will behave in a deployed
 environment.  To contribute to this project, make and test changes locally, or simply kick the tires before deciding to use it, there is a
  complete docker-compose stack defined in `docker/`, that can be operated with simple `make` commands.

### Prerequisites
In order to make use of the tooling described below, your system will require the following to be installed:
1. gnu make
1. Docker and docker-compose
1. go 1.14+
1. bash
1. git

### First Use
On first use of the docker-compose environment, several things need to take place.  The simplest path is to use `make`, running:
```
$ make init
```

behind the scenes, this downloads the latest postgres database container, boots it, and loads some starter data.  Then, the golang 1.14 container
 is downloaded to test and build the linux binary of this project.  Finally, the Alpine Linux container is downloaded, the binary gets packaged into
   it and started up, and configured to use the prepared postgres database.  At this point, the service is reachable at http://localhost:8081/api
   /v1/{...}

### Normal Startup/Shutdown
For pedestrian use cases, you might just want to start and stop the stack (this only works after a `make init` has been performed):

Starting:
```
make up
```

Stopping:
```
make down
```

### Redeploying the API 
During the course of development, frequently you'll want to redeploy a new build of the api for review and testing.  Again, the simplest way to do
 this is to use the make target:
```
make api
```

This will *_only_* build a new api service container and replace the running one with this new one.

### Completely Refresh the Running Stack
This is the factory reset button.  Sometimes you just want to burn it down and start over.  Maybe the database has accumulated a bunch of cruft.  Or
 you messed up a migration and hosed the schema.  Fear not:

```
make refresh
```

### The Nuclear Waste Removal Option
If you just want it gone:
```
make destroy
```

**NOTE**: This destroys the running docker-compose stack, associated data volume (e.g. the database), and orphans.  It does not delete the source
 images from docker.
