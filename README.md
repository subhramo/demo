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

Detailed Instructions:

1. Create a AWS EC2 machine using AWS console. 

2. Create a project directory under /home/ec2-user/projects & copy all the program files along with Dockerfile in the directory /home/ec2-user/projects

3. Install git using the below command:
      
       3.1 sudo yum install git -y

4. Use the below set of commands to establish a connectivity between your EC2 machine & remote github repo. The below instructions are referring to my personal github account. You should change the github repo details based on your solution.

        4.1 git init
        4.2 git add .
        4.3 git commit -m "first commit"
        4.4 git remote add origin https://github.com/subhramo/demo.git
        5.5 git push -u origin master
        
