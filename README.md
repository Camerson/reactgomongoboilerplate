# React/Redux, Go, MongoDB Boilerplate

This is a boilerplate project that includes a React/Redux application, Golang API, and MongoDB . 

#### This boilerplate integrates the following libraries for the frontend:
- React Router
- React Helmet
- Redux
- React Redux
- Redux Persist
- Axios (for fetching API endpoints)

#### This boilerplate includes the following functionalities on the frontend:
1. Layouts for different pages.
    1. For example, if you want different layouts for website and dashboard, 
    you can break the pages into different layouts for different components.
    2. Breaks react project into pages
2. A react router integration with:
    1. Exact page matches
    2. Dynamic pages
    3. Catch all 404 
3. A React Redux integration with:
    1. An implementation of mapstatetoprops and mapdispatchtoprops
    2. Persisted redux state with redux persist (default localStorage)
3. A React Helmet integration with:
    1. Loading head information from .env
    2. Loading head information from variables
    3. Loading head information from state
4. Basic form functionality:
    1. Changing input
        1. Text input
        2. Checkbox input
        3. Toggle show password
    2. Submitting form
        1. On Contact Page, this does not submit to database
        2. Todo table on home page submits to MongoDB   
    3. Show success message that hides after 3 seconds after form is submitted

#### This boilerplate includes the following Golang libraries
- Gorilla Mux
- Go MongoDB Driver

#### This boilerplate includes the following functionalities in the backend:
1. CRUD API routes to MongoDB
2. Integration to serve the react build when starting the entire app


## Instructions
### For Development
- cd frontend & npm start
- In another terminal: go run main.go
- Frontend will be on port 3000
- Backend will be on port 8080

### For Production
- cd frontend & npm run build
- go build .
#### If MacOS or Linux:
- ./reactgoboilerplate
#### If Windows:
- reactgoboilerplate.exe

##### Entire app will be on port 8080 in production

