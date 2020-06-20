node {
    def app

    stage('Clone repository') {
        checkout scm
        sh "git rev-parse --short HEAD > .git/commit-id"                        
        commit_id = readFile('.git/commit-id').trim()
    }

    stage('Build Image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("subhramo/demo")
    }

    stage('Test Image') {
        echo "Test Successful" 
        }
    }

    stage('Push Image') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
        docker.withRegistry('https://registry.hub.docker.com', 'subhramo') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
        }
    }
}
