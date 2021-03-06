# What is this Project about?

This is a demo project to illustrate dockerization of an app designed in golang & apply CI/CD methodologies to keep updating the code base, build & push container images through automation. This CI/CD process has been achieved through Jenkins pipeline. Once the new image has been built, it will be pushed into a Repo (either in Docker Hub or AWS ECR).

The Go Application is an extension of the original App used in our first Test Case: https://github.com/xUnholy/technical-tests. It exposes API endpoints like /version, /go, /helloworld etc. & show outputs based on the configurations.

On a high level, below tasks are involved in this project.
  
    1. Create an AWS EC2 Linux machine
    2. Install Go & Docker 
    3. Copy all project files including Dockerfile into a given folder
    4. Run the Program locally & conduct Unit Test
    5. Build & Run the Dockerfile to validate the application
    6. Install git & establish github integration with the code repo 
    7. Push/Pull all the project files into the Github repo
    8. Install Jenkins on AWS EC2 machine
    9. Integrate Jenkins with Github using webhook
    10. Design CI/CD Pileline using Jenkinsfile
    11. Create an AWS ECR or a Docker Hub Repo
    12. Integrate Jenkins with AWS ECR
    13. Finally, run the Jenkins Pipeline
    14. Additional Considerations
    
# The Architecture
  <img width="781" alt="Architecture" src="https://user-images.githubusercontent.com/30802518/85263165-1d7b0000-b4b2-11ea-9c78-abddea94155f.png">

# Detailed Instructions:

# 1. Create an AWS EC2 Linux machine. 
Using AWS CLI or Management console, create an EC2 Linux machine. Configure the Security Group for port 8080 (Jenkins) and port 8000 (Golong app)

# 2. Install Go & Docker

      2.1 sudo yum install go -y
      2.2 sudo yum install docker -y

# 3. Create a project directory
Create a directory under /home/ec2-user/projects & copy all the program files along with Dockerfile & Jenkinsfile in it. 

# 4. Build Program locally & conduct Unit Test
Use the below commands to compile & execute the program. Then, validate the output of the microservice & conduct an Unit Test.

       4.1 go build -o Demo -ldflags "-X main.gitCommit=$(git rev-list -1 HEAD)" .
       
       4.2 Execute: ./Demo
       
       * Output: starting http server
       * Open a browser: http://<IP>:8000/version
       * Output:
       myapplication: [  {version: 1.0,
                         lastcommitsha: d93d307161bc949160b5563772423807eceab7f1,
                         description : pre-interview technical test
                     } ]
       
       
       4.3 Run the Unit Test using: go test -v
           
           Sample Output if successful:
           === RUN   TestSum
           --- PASS: TestSum (0.00s)
           PASS
           ok  	golang-test	0.005s

# 5. Build & Run the Dockerfile
Build & test your docker images locally before automating the same using Jenkins pipeline in the subsequent steps.

       5.1 docker build -t demo/app .
       5.2 docker run -it -p 8000:8000 <DockerImageID>
       5.3 Open a browser: http://<IP>:8000/version
       5.4 Follow same steps as mentioned in step 4.2
       
# 6. Install git on AWS EC2 machine:
Use the below command to install git on your AWS EC2 machine so that you can pull/push your code changes into the Github Repository accordingly.

       6.1 sudo yum install git -y

# 7. Establish connectivity between AWS EC2 & Github repo. 
The below instructions are referring to my personal github account. You should change the github repo details based on your solution.

        7.1 git init
        7.2 git add .
        7.3 git commit -m "first commit"
        7.4 git remote add origin https://github.com/subhramo/demo.git
        7.5 git push -u origin master
        
# 8. Install Jenkins on the EC2 Linux machine  
Install Jenkins along with the recommended plugins like Github integration, Docker Pipeline etc.

        8.1 Use the instructions mentioned in the below link to install Jenkins on your machine.
   https://github.com/miztiik/DevOps-Demos/tree/master/setup-jenkins
        
        8.2 Give 'ec2-user' necessary permission to run the Jenkins Pipeline using the below commands:
            
            8.2.1 vi /etc/sysconfig/jenkins
            
            8.2.2 Find this $JENKINS_USER and change to “ec2-user”:
                  $JENKINS_USER="ec2-user"
            
            8.2.3 Then change the ownership of Jenkins home, webroot and logs:
                  chown -R ec2-user:ec2-user /var/lib/jenkins
                  chown -R ec2-user:ec2-user /var/cache/jenkins
                  chown -R ec2-user:ec2-user /var/log/jenkins
            
            8.2.4 Restart Jenkins and check the user has been changed:
                  service jenkins restart
                  ps -ef | grep jenkins
                  
# 9. Integrate Jenkins with Github using webhook.
Use the below instructions to integrate Jenkins to Github using Webhook which essentially ensures that when a commit has been made on a given branch, it should trigger a Jenkins job (SCM Polling) to perform configured tasks.

         9.1 In your Github account, go to Settings.
         9.2 Select Webhooks in the left panel
         9.3 Click on 'Add Webhook'
         9.4 Add the entry: ‘http://<IP>:8080/github-webhook’
         
# 10. Design the CI/CD Pipeline using Jenkinsfile
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
# 11. Create an AWS ECR or a Docker Hub Repo
From the AWS Management console, select Amazon Elastic Container Registry (ECR) and create a new Repo. Enable the option for Scan on Push which allow the pushed images to undergo vulnerability scan & security loop holes within the image using the CVE database.

# 12. Integrate Jenkins with ECR
Install aws cli on your AWS EC2 machine and use the below commands to manully verify if you are able to push any docker images into the newly created ECR Repo. The 3 steps mentioned below are to get the authentication token, tag the docker image & push the same into the ECR.

        12.1 aws ecr get-login-password --region ap-southeast-2 | docker login --username AWS --password-stdin <AWS Account ID>.dkr.ecr.ap-southeast-2.amazonaws.com
        
        12.2 docker tag demo:latest <AWS Account ID>.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest
        
        12.3 docker push <AWS Account ID>.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest
        
# 13. Finally, run the Jenkins Pipeline
At the Github repo, make changes in your code or Readme file & commit the same. It should trigger the Jenkins pipeline, build the docker image & push the same into the ECR.

# 14. Additional Considerations
While this given demo is a fairly simple illustration of a CI/CD pipeline, there are some additional steps which, if included in the design, would have been a complete solution. For example, if we can include SonarQube in one of the early stages of the pipeline, it would have scanned the vulnerabilities of the static codes along with its security loop holes even before compiling the program.

Similarly, with Clair Scanner, we can check the vulnerabilities of the docker images before pushing the same into the ECR. However, as part of this solution, we have already enabled ‘Scan on Push’ to check for vulnerabilities based on inbuilt CVE database in the ECR Repository.

# Refined Architecture
<img width="1001" alt="RF1" src="https://user-images.githubusercontent.com/30802518/85365955-b0c33c80-b569-11ea-9ed3-50ce4765de93.png">
