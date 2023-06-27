<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Vectorizeme</h3>

  <p align="center">
    A simple image and text vectorization service
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template/issues">Report Bug</a>
    ·
    <a href="https://github.com/othneildrew/Best-README-Template/issues">Request Feature</a>
  </p>
</div>

## About The Project

This project converts images or texts into embeddings on the same space using the OpenAI CLIP model. This model is hosted using a FastAPI server.

## Getting Started

### Run the server

1. Install the PIP packages
   ```sh
   pip install -r requirements.txt
   ```
2. Run the server (development)
   ```sh
   uvicorn src.main:app --reload
   ```

### Run the server with Docker

1. Build docker image

   ```sh
   docker build -t <name>/<image-name> .
   ```

2. Run docker image
   ```sh
   docker run -d -p <port>:<port> <name>/<image-name>
   ```

## Usage

_Please refer to the [Swagger Documentation](https://localhost:8000/swagger) once the server is running_

<!-- ROADMAP -->

## Roadmap

- [] Add Swagger docs
- [] Show image size and dimensions
- [] Show algorithm
- [] Add timestamp
- [] Refactor change Content model name to Vector

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<!-- CONTACT -->

## Contact

[@DumbIntuition](https://twitter.com/DumbIntuition)

<!-- Project Link: [https://github.com/KKogaa/vectorizeme](https://github.com/KKogaa/vectorizeme) -->

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[product-screenshot]: images/screenshot.png
[FastAPI.com]: https://img.shields.io/badge/FastAPI-005571?style=for-the-badge&logo=fastapi

<!-- [Go]:  -->
