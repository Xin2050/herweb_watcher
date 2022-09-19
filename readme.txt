




#before the program
0. rename config.dev to config.json

1. add  'alias chrome="/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome"'
to your .bash_profile, then reload the .bash_profile

2. close all the chrome processes

3. run chrome --remote-debugging-port=9222

4. copy the ws url to config.json, like [DevTools listening on ws://127.0.0.1:9222/devtools/browser/6a2a9afa-47ef-418e-8284-d27fb1ea717a]

5. set up your own mail server info to config.json and use mailer_test.go to test it

6. run cmd/rod.go