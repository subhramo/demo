This is a demo project to illustrate dockerization of an app designed in golang & apply CI methodolies to keep updating the code base & build activites through automation. This CI/CD process has been achieved through Jenkins pipeline. Once the new image has been built, it will be pushed into a Repo (either in Docker Hub or AWS ECR).

On a high level, below tasks are involved in this process.
  
    1. Created a AWS EC2 machine.
    2. Copied all project files including Dockerfile into a given folder.
    3. Installed git & established github integration with the code repo - subhramo/demo (Security Credentials)
    4. Pushed all the project files into the Github repo - subhramo/demo
    5. Installed Jenkins on AWS EC2 machine
    6. Integrated Jenkins with Github using webhook. (SCM Polling)
    7. Deigned the Pipe Jobs using Jenkins pipeline (Jenkinsfile)
    8. Integrated Jenkins with DockerHub or ECR (Security Credentials)
    
