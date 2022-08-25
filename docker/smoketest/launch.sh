#!/usr/bin/env sh

dockerize -wait tcp://${SMOKETEST_RECEIVER_URL#https://} -wait tcp://${SMOKETEST_RESULTS_URL#https://} -wait tcp://${SMOKETEST_CALLBACK_URL#https://} -wait-retry-interval 15s -timeout 5m

exec /ketch/bin/smoke.test
