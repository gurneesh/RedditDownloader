from urllib.request import urlopen, urlretrieve, Request
from json import JSONDecoder
import datetime
import os
import subprocess
import time
import datetime
import youtube_dl
import io
import re
import sys

class RedditDownloader:
    def __init__(self, subr):
        self.subr = subr
        self.url = 'https://www.reddit.com/r/' + self.subr + '.json'

    def req(self, previd = ''): 
        reddit_url = '%s?after=t3_%s' % (self.url, previd)
        print(reddit_url)
        rq = Request(reddit_url)
        rq.add_header('User-agent', 'Stylesheet images downloader Py3 v1')
        source = urlopen(rq).read()
        json = source.decode('utf-8')
        data = JSONDecoder().decode(json)
        return data

    def extract_imgur_album_urls(self, album_url):
        match = re.match("(https?)\:\/\/(www\.)?(?:m\.)?imgur\.com/(a|gallery)/([a-zA-Z0-9]+)(#[0-9]+)?", album_url)
        key = match.group(4)
        url = "https://imgur.com/a/" + key + "/layout/blog"
        response = urlopen(url)
        source = response.read().decode('utf-8')
        imageIDs = re.findall('.*?{"hash":"([a-zA-Z0-9]+)".*?"ext":"(\.[a-zA-Z0-9]+)".*?', source)
        image_urls = []
        for id in imageIDs:
            image_urls.append('https://imgur.com/' + id[0] + id[1])
        return (image_urls)

    def ydownload(self, link, foldername): #foldername = '/home/gurneesh/Downloads/' + url[25: -5]
        print(foldername)
        ydl_opts = {'outtmpl': foldername+'/%(title)s.%(ext)s',}
        with youtube_dl.YoutubeDL(ydl_opts) as ydl:
            ydl.download([link])

    def download(self, url, name):
        if os.path.exists(name):
            return
        print('Attempting to download {}'.format(url))
        filedata = urlopen(url).read()
        filehandle = open(name, 'wb')
        filehandle.write(filedata)
        filehandle.close()
        print('Successfully downloaded {}'.format(name))

            

if __name__ == '__main__':
    args = sys.argv
    downloader = RedditDownloader(args[1])
    #print(p[len(p) - 1].get('id'))
    FINISHED = False
    items = []
    foldername = '../Downloads/' + args[1]
    if not os.path.exists(foldername):
        os.mkdir(foldername)
    name = None
    temp = None
    previd = None
    count = 0
    while not FINISHED:
        name = None
        if (len(items) == 0):
            data = downloader.req()
            items = [x['data'] for x in data['data']['children']]
            # print(len(items))
            previd = items[len(items) - 1].get('id')
        else :
            temp = items[len(items) - 1].get('id')
            data = downloader.req(previd)
            items = [x['data'] for x in data['data']['children']]
        if (len(items) > 1):
            previd = items[len(items) - 1].get('id')
        for item in items:
            try: #time.sleep(4)
                print(item.get('url'))
                if item.get('url'):
                    link = item.get('url')
                    count+=1
                    id = item.get('id')
                    if link.endswith('.jpg'):
                        name = foldername + '/' + id + '.jpg'
                        downloader.download(link, name)
                    elif 'imgur.com/a/' in link or 'imgur.com/gallery/' in link:
                        u = downloader.extract_imgur_album_urls(link)
                        for i in u:
                            print('imgur_url', i)
                            name = foldername + '/' + i[-11: ]
                            downloader.download(i, name)
                    elif link.endswith('.gif'):
                        name = foldername + '/' + id + '.gif'
                        downloader.download(link, name)
                    elif link.endswith('.mp4'):
                        name = foldername + '/' + id + '.mp4'
                        downloader.download(link, name)
                    elif link.endswith('.webm'):
                        name = foldername + '/' + id + '.webm'
                        downloader.download(link, name)
                    elif link.endswith('.png'):
                        name = foldername + '/' + id + '.png'
                        downloader.download(link, name)
                    elif 'gfycat' in link:
                        downloader.ydownload(link, foldername)
                    else :
                        downloader.ydownload(link, foldername)
            except Exception as e:
                print(e)
        if (temp == previd):
            FINISHED = True
