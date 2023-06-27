from typing import List


class Content:
    url: str
    vector: List[float]
    text: str
    algorithm: str

    def __init__(
        self,
        url: str = None,
        vector: List[float] = None,
        text: str = None,
        algorithm: str = None,
    ):
        self.url = url
        self.vector = vector
        self.text = text
        self.algorithm = algorithm

    def get_filename(self) -> str:
        if self.url:
            filename = self.url.split("/")[-1]
            return filename
        else:
            raise ValueError("URL is not provided")

    def get_file_extension(self) -> str:
        filename = self.get_filename()
        if "." in filename:
            file_extension = filename.split(".")[-1]
            return file_extension
        else:
            raise ValueError("File extension is not found")
