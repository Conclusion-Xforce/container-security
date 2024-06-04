# Do better

The goal of this exercise is to improve the security and quality of the provided containerized application by implementing good practices and security measures. Focus on the Dockerfile, even though I'm sure the Go code can be improved too. Use common sense and the good practices we have discussed in class.

For reference: a build of the provided Dockerfile is available at `quay.io/conclusionxforce/bad-container`.

Feel free to use any tools or resources you find useful. The final result should be a Dockerfile that builds a secure and efficient container image. Make sure to document your findings and the changes you made!

> Hint: use the `valkey/valkey:7.2` container image and a Kubernetes `Service` to create a standalone containerized Redis.

At the end of the exercise, both Redis/Valkey and the provided application should be running in OpenShift. Good luck!

## Resources

- [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)
- [Dockerfile best practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
- [Creating images | OpenShift Container Platform 4.14](https://docs.openshift.com/container-platform/4.14/openshift_images/create-images.html#images-create-guide-openshift_create-images)