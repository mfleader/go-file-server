from pathlib import Path
from typing import (BinaryIO, Generator)

import httpx


def upload_bytes(
    file: BinaryIO,
    # chunk_size: int = 262_144_000
    # chunk_size: int = 1024
    chunk_size: int = 1024 * 1024
) -> Generator[bytes, None, None]:
    # 250 MiB == 250 * 1024 * 1024 == 262_144_000
    contents = 'dummy'
    pointer = 0
    while len(contents):
        file.seek(pointer)
        pointer += chunk_size
        contents = file.read(chunk_size)
        yield contents

import requests
from requests_toolbelt.streaming_iterator import StreamingIterator


def post_file(url: str, filepath: Path):
    size = filepath.stat().st_size
    # r = httpx.post(
    #   url,
    #   # files = {
    #   #   'document': open(filepath, 'rb')
    #   # }
    #   data=upload_bytes(filepath.open('rb'))
    #   # data=b'big-data'
    # )
    r = requests.post(
      url,
      data=StreamingIterator(size, upload_bytes(filepath.open('rb'))),
      stream=True
    )
    print(r.request)
    print(r.text)
    return r


def main():
  post_file(
    'http://localhost:4000/stream',    
    Path('/home/mleader/50K.txt')
    # Path('/home/mleader/Pictures/bayes-ab.png')
    # Path('/home/mleader/lg.txt')
    # Path('/home/mleader/10G.txt')
  )


if __name__ == '__main__':
  main()