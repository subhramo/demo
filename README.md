This is a demo project to illustrate dockerization of an app designed in golang & apply CI methodolies to keep updating the code base & build activites through automation. This CI/CD process has been achieved through Jenkins pipeline. Once the new image has been built, it will be pushed into a Repo (either in Docker Hub or AWS ECR).

On a high level, below tasks are involved in this process.
  
    1. Create a AWS EC2 machine.
    2. Copy all project files including Dockerfile into a given folder
    3. Run the Program locally & conduct Unit Test
    4. Build & Run the Dockerfile to validate the application
    5. Install git & established github integration with the code repo - subhramo/demo (Security Credentials)
    6. Pushed all the project files into the Github repo - subhramo/demo
    7. Installed Jenkins on AWS EC2 machine along with recommended plugins & Docker Pileline
    8. Integrated Jenkins with Github using webhook. (SCM Polling)
    9. Deigned the Pipe Jobs using Jenkins pipeline (Jenkinsfile)
    10. Integrated Jenkins with DockerHub or ECR (Security Credentials)
    
    

## Detailed Instructions:

1. Create a AWS EC2 Linux machine using AWS console. Ensure to configure the security group for port 8080 (Jenkins) and port 8000 (Golong app)

2. Create a project directory under /home/ec2-user/projects & copy all the program files along with Dockerfile in the directory /home/ec2-user/projects

3. Build & Run the Dockerfile

       3.1 docker build -t demo/app .
       3.2 docker run -it -p 8000:8000 <DockerImageID>
       3.3 Open a browser: http://<IP>:8000/version
       
       Sample Output:
       
4. Install git using the below command:
      
       4.1 sudo yum install git -y

5. Use the below set of commands to establish a connectivity between your EC2 machine & remote github repo. The below instructions are referring to my personal github account. You should change the github repo details based on your solution.

        5.1 git init
        5.2 git add .
        5.3 git commit -m "first commit"
        5.4 git remote add origin https://github.com/subhramo/demo.git
        5.5 git push -u origin master
        
6. Install Jenkins on the EC2 Linux machine using the instructions mentioned in the link. Also install set the set of recommended plugins like Github integration, etc.

        6.1 https://github.com/miztiik/DevOps-Demos/tree/master/setup-jenkins
        
        6.2 Give 'ec2-user' necessary permission to run the Jenkins Pipeline using the below commands:
            
            6.2.1 vi /etc/sysconfig/jenkins
            
            6.2.2 Find this $JENKINS_USER and change to “ec2-user”:
                  $JENKINS_USER="ec2-user"
            
            6.2.3 Then change the ownership of Jenkins home, webroot and logs:
                  chown -R ec2-user:ec2-user /var/lib/jenkins
                  chown -R ec2-user:ec2-user /var/cache/jenkins
                  chown -R ec2-user:ec2-user /var/log/jenkins
            
            6.2.4 Restart Jenkins and check the user has been changed:
                  service jenkins restart
                  ps -ef | grep jenkins
                  
7. Integrated Jenkins with Github using webhook.

         7.1 In your Github account, go to Settings.
         7.2 Select Webhooks in the left panel
         7.3 Click on 'Add Webhook'
         7.4 Add the entry: ‘http://<IP>:8080/github-webhook’
