node {
    def app
    def commit_id

    stage('Clone repository') {
        checkout scm
        sh "git rev-parse --short HEAD > .git/commit-id"                        
        commit_id = readFile('.git/commit-id').trim()
    }
 
    stage('Unit Test') {
        sh "go test -v"
        }
    
    stage('Build image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("subhramo/demo")
    }

    stage('Push image') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
        sh "aws ecr get-login-password --region ap-southeast-2 | docker login --username AWS --password-stdin 825030697311.dkr.ecr.ap-southeast-2.amazonaws.com"
        sh "docker tag demo:latest 825030697311.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest"
        sh "docker push 825030697311.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest"
        /*docker.withRegistry('https://registry.hub.docker.com', 'subhramo') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
        }*/
    }
}
