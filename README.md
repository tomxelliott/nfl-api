# nfl-api
Tester API built with Go using NFL teams as test specimen

## Building the app
All of the build instructions assume you have cloned the repo and are running commands from the root directory.

### Building the Docker Container
Since bundling the Go runtime into a Docker container can lead to very bloated containers, I have opted for scratch which leads to lightweight Docker images whilst still delivering all that we need.

To build locally:

```
$ docker build -t <your_choice_of_name> .
```

You can then run the image locally:

```
$ docker run -p 1234:5000 <your_choice_of_name>
```

If you navigate `localhost:1234` in your browser you will be able to see the initial data. 

### Pulling the container from Docker Hub
If you would rather just try this little API out without cloning the code, you can pull the image from Docker Hub

```
docker pull tomxelliott/nfl-api-tom
```

### Deploying to Elastic Beanstalk
Using the Elastic Beanstalk cli you can simply run
```
$ eb init
$ eb create
```
Default values will be sufficient to get the application running on Elastic Beanstalk.

The URL for your own Elastic Beanstalk deployment can be found on the AWS console.


## Making API calls
### Post to the API
```
$ curl -H "Content-Type: application/json" -d '{"Name": "Dallas Cowboys", "Info": {"Coach": "Jason Garrett", "NoOfTitles": 5, "Stadium": "AT&T Stadium"}}' http://localhost:1234/add-team
```

Since this was only a quick example, there is no persistence of data
