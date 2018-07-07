#!/usr/bin/python3

from urllib.request import urlopen, urlretrieve, Request
from json import JSONDecoder
import os
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
        try:
            rq = Request(reddit_url)
            rq.add_header('User-agent', 'Stylesheet images downloader Py3 v1')
            source = urlopen(rq).read()
        except Exception as e:
            print(e)
            return e
        json = source.decode('utf-8')
        data = JSONDecoder().decode(json)
        return data

    #method to extract urls of all the images of photos from imgur ablums
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

    #method to download video content using youtube-dl module
    def ydownload(self, link, foldername): #foldername = '/home/gurneesh/Downloads/' + url[25: -5]
        print(foldername)
        ydl_opts = None
        if 'gfycat' in link:
            ydl_opts = {'outtmpl': foldername+'/%(webpage_url_basename)s.%(ext)s',}
        else:
            ydl_opts = {'outtmpl': foldername+'/%(title)s.%(ext)s',}
        with youtube_dl.YoutubeDL(ydl_opts) as ydl:
            ydl.download([link])

    #method to download any mime type whose url is directly available in the json of the page
    def download(self, url, name):
        if os.path.exists(name):
            print('Already Downloaded ')
            return
        print('Attempting to download {}'.format(url))
        filedata = urlopen(url).read()
        filehandle = open(name, 'wb')
        filehandle.write(filedata)
        filehandle.close()
        print('Successfully downloaded {}'.format(name))


#main function to use all the methods of the class to scrape a subreddit.
#TO DO reduce the code into methods of the above class.
def main(subr):
    downloader = RedditDownloader(subr)
    FINISHED = False
    items = []
    #if os.path.exists(os.getcwd()+'/'+subr):
    #    print('downloaded')
    #    return
    foldername = os.getcwd()+'/'+subr
    if not os.path.exists(foldername):
        os.mkdir(foldername)
    name = None
    temp = None
    previd = None
    while not FINISHED:
        name = None
        if (len(items) == 0):
            data = downloader.req()
            if(type(data)==dict):
                items = [x['data'] for x in data['data']['children']]
                previd = items[len(items) - 1].get('id')
            else:
                return
        else :
            temp = items[len(items) - 1].get('id')
            data = downloader.req(previd)
            items = [x['data'] for x in data['data']['children']]
        if (len(items) > 1):
            previd = items[len(items) - 1].get('id')
        for item in items:
            try:
                print(item.get('url'))
                if item.get('url'):
                    link = item.get('url')
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
            print('Completly scrapped {}'.format(subr))
            FINISHED = True

if __name__=='__main__':
    args = sys.argv
    main(args[1])
