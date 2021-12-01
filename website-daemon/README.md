# Test

1.  Run `renew-certificate.sh`.
    ```
    $ ./renew-certificate.sh [-i certificate_download_url] [-o certificate_path]
    ```

2. Validate if `renew-certificate.sh` downloads certificate to specified path.
    
    Example:
    ```
    $ ls /home/nobellet/short-lived-cert/website-daemon/certificate.pdf
    /home/nobellet/short-lived-cert/website-daemon/certificate.pdf
    ```

# Example to set up cron job

1. Edit crontab file.
    ```
    $ crontab -e
    ```

2. Schedule `renew-certificate.sh` to run every day by adding the following to the crontab file opened in Step 1.
    
    Example:
    ```
    0 0 * * * /home/nobellet/short-lived-cert/website-daemon/renew-certificate.sh
    ```

3. Save and exit crontab file.