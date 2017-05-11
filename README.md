# S3 Utility

This a utility written in Go that uses Amazon's Go SDK to interact with the
AWS S3 service.  Amazon provides a set of command line utilities, but they're
written in Python and require external Python libraries to run.

# Quickstart

    go get -u github.com/sh3rp/s3

# Usage

The utility uses four environment variables:

    S3_ACCESS_KEY (required)
    S3_SECRET_ACCESS_KEY (required)
    S3_BUCKET
    S3_REGION

Each are what you would expect, but the first two are required.  The latter
two can be supplied on the command line.

Run 's3 help' for a list of commands you can run.  Essentially, in a command
line such as 's3 ls', you can either not supply a command and get a full
listing of the bucket or if you supply an argument, that will be used as the
filter criteria.
