---
description: How to backup your photos and videos in Photos+ on Storj DCS.
---

# Photos+ Integration Guide

## Introduction

**Photos+** is an iOS/Android App. It allows you to store and manage your photos and videos in your own cloud storage account. You can use **Photos+** with **Storj DCS** S3-compatible storage.

**Main site**: [https://photosplus.app/](https://photosplus.app/)

Mobile app download links:

**Apple App Store**: [https://apps.apple.com/us/app/photos-cloud-library/id1310744251](https://apps.apple.com/us/app/photos-cloud-library/id1310744251)

**Google Play Store**: [https://play.google.com/store/apps/details?id=com.pixegram.photosplus](https://play.google.com/store/apps/details?id=com.pixegram.photosplus)

## Configure Photos+ to use Storj DCS

1. When starting **Photos+**, you will be shown an empty Library.

![](https://lh4.googleusercontent.com/HuatGXoi7xhlo6Q9Y07aoA\_MbSDRnr\_3vAn0zR5\_WNuThINH\_q9JKH64gG6QFxLBEp6oto1M8mMrTiOvWU12vLhyW7XDwgHNO7-Oqx9qydRefti6NCt7eLPvvbseUIhPsuA8AXWsohqgqtWXEg)

2\. Click on the person-shaped icon in the lower right corner to show your Profile page.

![](https://lh5.googleusercontent.com/ppWLkNtjwtC7xBBtCsLHKusFd7KRLAt0xDUObOI3-EF7yS7s3JnNvUqB1LHmZDotPDgNQZVYkG1HHhRL7Z95JNyfnLwgS0jTckLOAQrRkPS\_S4wmRwlKTL4-UeqHEFDMOzIIZEDZ2CSNNSKV7Q)

3\. Click on the cog wheel icon on the upper right to access Settings.

![](<../.gitbook/assets/IMG\_EDAEA36C758C-1 (2).jpeg>)

4\. Click **Configure** to begin setting up Storj DCS for Photos+.

![](../.gitbook/assets/IMG\_F65448A9A594-1.jpeg)

5\. Select **S3 Storage** and enter “gateway.storjshare.io” as the URL, along with your corresponding access key and secret key. See [Generate Gateway Credentials](../getting-started/gateway-mt/#generate-credentials-to-the-gateway-mt) for details on creating these access credentials.

![](../.gitbook/assets/IMG\_419569E7945F-1.jpeg)

6\. Once you’ve completed configuring Storj DCS as your Cloud Storage Provider, to see Storj in action, you’ll need to click the icon in the lower right corner composed of four squares to return to your Library. From there, click the three-dot icon in the upper right corner.

![](https://lh4.googleusercontent.com/E7V5bWlQl1F0kZaI7w10OIWcRHDt4uREPKvQ6O3sqot3OI2g4l-1Qmo41TipQ9HRSY51OZdRd3hcAaqUb\_ESSC5rY8iiMABlR9ttOUPwkraeZfxIDg-RSE27YlfUUcnhBwjbbk8ksVPiIir2RA)

7\. Click **Add Album** and type in your chosen Album name. Then click the large checkbox in the center of the screen to continue.

![](https://lh4.googleusercontent.com/\_tsYVh7e80hrjuTu\_U\_P92hyI1JQJloDgPVpOJbGUaHc\_aj4XbNpl6zUJdkpZshWo7ENEW5v0akaSqGGRfSk1jbh0ID1s3nQzNpsK2lV\_6DltpfW5fkorfg4\_\_U9oALzHottV7I0Svm1mY4n3w)

8\. Once your album is created, click its name to view it. While viewing the album, you can add content to it using the plus-sign icon in the upper right corner. This will bring up the “Select Album” screen. Click the album from which you want to add photographs.

![](https://lh3.googleusercontent.com/F0OhNS\_gS9UAvZU1M3C2N9-t2LpImtZFOxzwmpcWcJcg\_XG6kDLOzx7eFfkFFRhBkqyyEIe2J47CiTPNN1kCj2hDp0K2Zuyz6tKWhZJ4nQqk\_FIJvZXSgmh3MIiVSL2QN0I7HLleo--lrDLHfA)

9\. When viewing an album, select individual media to add by clicking each item, or use the blank checkbox next to the **Add** checkbox in the upper right to add all items. Once you’ve selected all the items your want in your album, click **Add** to continue.

![](https://lh6.googleusercontent.com/WOWRxlRKQFrXvlnZGTyHpySngtdQfkQojxtXTU4Z3eNYfgD1GkEfqdpplqSQJfEpGW8nRvh8yOKjLx0GUx0yo9AEQ1ip\_s4pESX16MasjK9DPNprCw8LvGWCB0-PuIcC0jV881fMfpIky2ulRw)

10\. At this point you should see an album in your Library. You’re ready to synchronize your media to Storj. Click the circular arrow icon at the bottom center of the screen.

![](https://lh5.googleusercontent.com/OpiE-flzbWASQ\_xmfIRzilYvVj0If-wblNwe6M3uCHhDjhKVrw9GyHFSg3vMJGGjwtoQJvNyHzeCVvwXNHJqTGUoEqAvye5C39rfIeIkGrLqR6AOiFqC3X7oDCmedxtKhICnmu0kCqTcxNCmlQ)

11\. Click **Start** to begin synchronizing your library. You may check via the Storj DCS dashboard or use any of our command line tools to see your library has been backed up to Storj.

![](https://lh4.googleusercontent.com/QwQtZYzE0sYPyEa2nwxaoQI4ofUSihdl5QHjRkXE2JrdlFSasYpb0935sRvG48WI9jMF\_zjhDquHS1spOkWHqMq4TQehCLzkACARqzh4UkNcIM6OUiAveNQjtAEcSYS4cWlgPewFSqNADcxyvw)

## How to automatically import your photos from Google Photos to Photos+

1. Request a Google Takeout of your Google Photos by visiting [https://takeout.google.com](https://takeout.google.com/)
2. Select Google Photos only for export.

![](<../.gitbook/assets/image (62).png>)

3\. Choose your export location. If you have enough free storage space on Google Drive or Dropbox, select it as the location and the 50GB export size. If you don’t have enough free storage, select “Send download link via email” and the 10GB export size. Please wait until you get an email from Google that your export has been completed.

![](<../.gitbook/assets/image (34).png>)

4\. If you have not done so already, [sign up for Photos+](photos+-integration-guide.md#introduction) and [configure your Storj account through the Photos+ iOS/Android app](photos+-integration-guide.md#configure-photos+-to-use-storj-dcs). Then sign in to your Photos+ account on your **desktop or tablet** browser at [https://photosplus.app](https://photosplus.app/).

5\. Click **Import** from the main menu

6\. Click **Select from Google Drive**, **Dropbox** or **Select from Computer** depending on where you saved your Google Takeout file.

![](<../.gitbook/assets/image (32).png>)

7\. Once our system downloads your file from Google Drive or Dropbox, or it receives it via the upload from your computer, it will parse it and show you the albums and the number of photos and videos it contains.

8\. Confirm the selections and click **Start Import** to begin the import process.

![](<../.gitbook/assets/image (184).png>)

Your photos and videos will automatically appear in the Library section as they are imported, and you may track the import progress by going back to the Import section.
