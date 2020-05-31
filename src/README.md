# Streetlity Driver
Streetlity Driver provides a platform for managing resources like image, video, documentation,... etc

# Methods

## Upload
Upload send files from client to the server for storaging. When a duplicated file is sent, depends on the params `wmethod`:
- `0`: ignore the uploading process
- `1`: override the existed file
- `2`: rename the uploading file

## Download
Download send files from server to the client, downloaded file must be existed on the server. If not, 


## Delete
Delete delete files on server.