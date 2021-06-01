# Adlive-Go

## Initialize Screen

1. Screen App start initializing which collects Mac address, type of devices,...
2. Connect to google-firebase to get its deviceToken which will be used to receive notification from server
3. Generate QR Code from above information.
4. User uses his/her mobile app to scan QR Code.
    
    4.1 If device exists in server then he/she needs to contact to screen's owner.
    
    4.2 If it does not, then user starts input screen's information such as screen name, location, area, etc. Then submit to server.

5. Server adds a new screen into database and notifies accessToken and refreshToken to the device via device token
6. The screen receives accessToken and refreshToken from the server.

## TODO:
1. Notify accessToken and refreshToken to devices via device token
2. Implement workers (machinery) to receive logs (activities) and counting played from devices
3. Implement video management (uploading, editing, listing)
4. Create a sample client in flutter to demo how to connect to server via GRPC
5. Create a schedule to manage campaign schedule. (eg: notify to devices to download videos)
6. Implement function to randomly playing video from weighted number for video.
7. Write unit test or integration test for all functions (or at least important function)
8. Investigate issue cannot modify data in SQLite (for testing purpose) (optional)

