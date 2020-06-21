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
         docker.withRegistry('https://825030697311.dkr.ecr.ap-southeast-2.amazonaws.com/demo', 'ecr:ap-southeast-2:subhra')
         app.push("latest")
        }
    }
}
