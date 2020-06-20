node {
   def commit_id
   stage('Preparation') {
     checkout scm
     sh "git rev-parse --short HEAD > .git/commit-id"                        
     commit_id = readFile('.git/commit-id').trim()
   }
   stage('test') {
     echo "test Successful"
     }
   stage('docker build/push') {
     docker.withRegistry('https://index.docker.io/v1/', 'subhramo') {
       def app = docker.build("subhramo/demo:${commit_id}", '.').push()
     }
   }
}
