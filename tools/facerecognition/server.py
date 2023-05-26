#!/usr/bin/env python3

import face_recognition.api as fra
import os
import json

from http.server import BaseHTTPRequestHandler, HTTPServer

known_names = list()
known_face_encodings = list()

class RequestHandler(BaseHTTPRequestHandler):
    def do_POST(self):
        
        content_length = int(self.headers['Content-Length'])
        post_data = self.rfile.read(content_length)
        data = json.loads(post_data.decode('utf-8'))
        try:
            file_dir = data["path"]
            print(file_dir)
            unknown_encodings = fra.face_encodings(fra.load_image_file("./ramdisk/"+file_dir))
            if len(unknown_encodings) == 0:
                str = "Face not detected"
            else:
                ans = list()
                for unknown_encoding in unknown_encodings:
                            
                    distances = fra.face_distance(known_face_encodings,unknown_encoding)
                    for index,val in enumerate(distances):
                        if val < 0.6:
                            ans.append(known_names[index])
                            break
                str = ','.join(ans)
        except:
            str = "Failed to decode"
        print(str)
          
        # 处理POST请求的数据
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        response = {'message': 'Received POST request', 'names': str}
        self.wfile.write(json.dumps(response).encode('utf-8'))

    def do_GET(self):
        # 在这里处理GET请求
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()
        response = '<html><body><h1>Hello, World!</h1></body></html>'
        self.wfile.write(response.encode('utf-8'))

def run(server_class=HTTPServer, handler_class=RequestHandler, port=43221):
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    print(f'Starting server on port {port}...')
    httpd.serve_forever()

if __name__ == '__main__':
    with open("config.json") as f:
        config = json.load(f)
    filesdir = [os.path.join(config["known_img_dir"],file) for file in os.listdir(config["known_img_dir"])]
    known_names = [os.path.splitext(os.path.basename(filedir))[0] for filedir in filesdir]
    known_face_encodings = [fra.face_encodings(fra.load_image_file(file))[0] for file in filesdir]
    print(known_names)
    run()

    
    
