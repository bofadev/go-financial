# go-financial
A golang backend project with react front end. Meant to demonstrate my understanding and ability to create fullstack web applications. Setup to deploy on heroku.

Currently hosted at: https://radiant-forest-56799.herokuapp.com/

## Intents
Design a golang RESTful backend service that
* Serves web requests
* Accepts data writes
* Provides data reads
* Houses a peristed database (memory/disk)
* Has clear interfacing between between database use and actual database implementation
  * Makes mocking functionality via some fake easy
  * Makes it possible to use any number of actual database implementations without affecting logic accessing the database
* Uses value/reference intentionally
  * Avoid data copying where desirable to
* Assume someone who will never meet me needs to work with the code
  * As self documenting as possible
  * Clear error logging
* Uses testing and test reporting
 * Test files alongside each internal package
 * Base dir script to run all internal package tests with ease

Deploys on heroku
* Conserve efforts to spend on the project work and not on hosting
* Reliable

Design a front end that
* Uses a clean and light weight frontend UI: Bulma
* Uses as few React features as necessary
* Uses as many out-of-the-box frontend features to provide input validation as possible
  * Disabling buttons, input types
  
Git
* Use a simple branching scheme
 * Work on dev
 * Deploy from deploy on Heroku
 * Merge stable milestones to master
 * Tag deploy branch for easy finding of deploy commit points
 * Tag master for easy finding of stable commit points

## Stages of development

Golang Pen & Paper

Golang Startup

Golang Testing Part 1

Golang Serving & Endpoints

Frontend Startup

Frontend Pen & Paper

Frontend Iteration: Blocking Out

Frontend Iteration: Fetching Data Read

Frontend Iteration: Input Validation

Frontend Iteration: Fetching Data Write & Redraw

Backend Iteration: JSON String Responses

Frontend UX bug searching

Golang Testing Part 2

Deploying To Heroku

Pushing to Git & Connecting To Heroku

Getting local PostgreSQL provider working

Setting up remote heroku PostgreSQL provider

Getting remote PostgreSQL provider working

## Wishlist
* Testing front end with Jasmine
* Research load testing
* Feature: User login
 * Email + Facebook + Google authentication (Using Auth0?)
 * Customized database tables per user
* Mobile friendly front end
* Heroku continuous integration connetion to golang tests
* Standardized backend logging rather than stdout printing
* Adhere main/internals to some less arbitrary coding standard
* Benchmark backend memory usage and latency
