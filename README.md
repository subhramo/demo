This is a demo project to illustrate dockerization of an app designed in golang & apply CI methodolies to keep updating the code base & build activites through automation. This CI/CD process has been achieved through Jenkins pipeline. Once the new image has been built, it will be pushed into a Repo (either in Docker Hub or AWS ECR).

On a high level, below tasks are involved in this process.
  
    1. Create a AWS EC2 machine.
    2. Copy all project files including Dockerfile into a given folder
    3. Run the Program locally & conduct Unit Test
    4. Build & Run the Dockerfile to validate the application
    5. Install git & established github integration with the code repo 
    6. Push/Pull all the project files into the Github repo (subhramo/demo)
    7. Install Jenkins on AWS EC2 machine along with recommended plugins along with Github Integration & Docker Pileline
    8. Integrated Jenkins with Github using webhook. (SCM Polling)
    9. Design CI/CD Pileline using Jenkinsfile
    10. Create an AWS ECR or a Docker Hub Repo
    11. Integrated Jenkins with ECR
    12. Finally, Run the Jenkins Pipeline
    

# Detailed Instructions:
------------------------
# 1. Create an AWS EC2 Linux machine. 
Using AWS CLI or Management console, create an EC2 Linux machine. Configure the Security Group for port 8080 (Jenkins) and port 8000 (Golong app)

# 2. Create a project directory
Create a directory under /home/ec2-user/projects & copy all the program files along with Dockerfile & Jenkinsfile in it. 

# 3. Build Program locally & conduct Unit Test
Using the below commands to compile & execute the program, validate the outout & conduct Unit Test.

       3.1 go build -o Demo -ldflags "-X main.gitCommit=$(git rev-list -1 HEAD)" .
       
       3.2 Execute: ./Demo
       
       * Output: starting http server
       * Open a browser: http://<IP>:8000/version
       * Output:
       myapplication: [  {version: 1.0,
                         lastcommitsha: d93d307161bc949160b5563772423807eceab7f1,
                         description : pre-interview technical test
                     } ]
       
       
       3.3 Run the Unit Test using: go test -v
           
           Sample Output if successful:
           === RUN   TestSum
           --- PASS: TestSum (0.00s)
           PASS
           ok  	golang-test	0.005s

# 4. Build & Run the Dockerfile
Build & test your docker images locally before automating the same using Jenkins pipeline in the subsequent steps.

       4.1 docker build -t demo/app .
       4.2 docker run -it -p 8000:8000 <DockerImageID>
       4.3 Open a browser: http://<IP>:8000/version
       4.4 Follow same steps as mentioned in step 3.2
       
# 5. Install git on AWS EC2 machine:
Use the below command to install git on your AWS EC2 machine so that you can pull/push your code level changes into the Repo accordingly.

       5.1 sudo yum install git -y

# 6. Establish connectivity between AWS EC2 & Github repo. 
The below instructions are referring to my personal github account. You should change the github repo details based on your solution.

        6.1 git init
        6.2 git add .
        6.3 git commit -m "first commit"
        6.4 git remote add origin https://github.com/subhramo/demo.git
        6.5 git push -u origin master
        
# 7. Install Jenkins on the EC2 Linux machine  
Install the recommended plugins along with Github integration, Docker Pipeline etc.

        7.1 https://github.com/miztiik/DevOps-Demos/tree/master/setup-jenkins
        
        7.2 Give 'ec2-user' necessary permission to run the Jenkins Pipeline using the below commands:
            
            7.2.1 vi /etc/sysconfig/jenkins
            
            7.2.2 Find this $JENKINS_USER and change to “ec2-user”:
                  $JENKINS_USER="ec2-user"
            
            7.2.3 Then change the ownership of Jenkins home, webroot and logs:
                  chown -R ec2-user:ec2-user /var/lib/jenkins
                  chown -R ec2-user:ec2-user /var/cache/jenkins
                  chown -R ec2-user:ec2-user /var/log/jenkins
            
            7.2.4 Restart Jenkins and check the user has been changed:
                  service jenkins restart
                  ps -ef | grep jenkins
                  
# 8. Integrated Jenkins with Github using webhook.
Use the below instructions to integrate Jenkins to Github using Webhook which essentially ensures that when a commit has been made on a given branch, it should trigger a Jenkins job (SCM Polling) to perform configured tasks.

         8.1 In your Github account, go to Settings.
         8.2 Select Webhooks in the left panel
         8.3 Click on 'Add Webhook'
         8.4 Add the entry: ‘http://<IP>:8080/github-webhook’
         
# 9. Design the CI/CD Pipeline using Jenkinsfile
The Jenkinsfile represents the skeleton structure of the CI/CD pipeline invoking the appropriate jobs in each stage.

        node {
         def app
    
            stage('Clone repository') {
                checkout scm
             }

            stage('Unit Test') {
             }

            stage('Build image') {
                  app = docker.build()
            }

            stage('Push image') {
                }
             }
# 10. Create an AWS ECR or a Docker Hub Repo
From the AWS Management console, select Amazon Elastic Container Registry (ECR) and create a new Repo. Enable the option for Scan on Push which allow the pushed images to undergo vulnerability scan & security loop holes within the image using the CVE database.

# 11. Integrated Jenkins with ECR
Install aws cli on your AWS EC2 machine and use the below commands to manully verify if you are able to push any docker images into the newly created ECR Repo. The 3 steps mentioned below are get the authentication token, tag the docker image & push the same into the ECR.

        1. aws ecr get-login-password --region ap-southeast-2 | docker login --username AWS --password-stdin <AWS Account ID>.dkr.ecr.ap-southeast-2.amazonaws.com
        
        2. docker tag demo:latest <AWS Account ID>.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest
        
        3. docker push <AWS Account ID>.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest
        
# 12. Finally, Run the Jenkins Pipeline
Make any change in your code or Read me file & make a commit on the Git Repo. It should trigger the Jenkins pipeline, build the docker image & push the same into the ECR.
