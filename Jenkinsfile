import hudson.model.*
import hudson.EnvVars
import groovy.json.JsonSlurperClassic
import groovy.json.JsonBuilder
import groovy.json.JsonOutput
import groovy.json.*
import java.net.URL
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
        docker.withRegistry('https://registry.hub.docker.com', 'subhramo')
        app.push()
        }
}
