Fest Slotter
=

## Run locally
`GOPATH=$PWD/bin` - ensures all downloaded packages are present within this project folder.
GOROOT=go/installed/dir

```
go build
```
will create a new file that's named after the module **festility**.


```
go run festility
# OR
go run <package-name>
```
which should start the server

## Run on Docker
Run the following in the root dir of the project
<pre>
# BUILD THE IMAGE
docker build -t festility .

# RUN THE CONTAINER
docker run -it -p 8080:8080 festility
</pre>

-Thanks from the future
https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/go-buildfile.html
https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/java-se-procfile.html
https://www.youtube.com/watch?v=5GINgmS93Mc

## MongoDB docker
```
docker run -d -p 27017:27017 mongo
```
