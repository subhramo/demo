node {
    def app
    
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
        sh "/var/lib/jenkins/workspace/Final/push.sh"
        }
}
