## gcpsqlclonner
### Setup
two system varible need to be set 

```GOOGLE_APPLICATION_CREDENTIALS``` ---  SA key

```GOOGLE_PARENT``` may set for project to org level.
 ```projects/{projectnumber}```
or 
```organizations/{org_number}```

### API endpoints

```/api/v1/csqlall``` -- get all sql instaqnces from gcp assets manager

```/api/v1/csql/{project}``` --- get all sql instances from project

```/api/v1/csql/{project}/{instance}/clone ``` --- clone instace within project 

```/api/health``` --- health endpoint
