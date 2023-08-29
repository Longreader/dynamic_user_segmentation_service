<a name="readme-top"></a>


<h1 align="center"><b>Dynamic user segmentation service<b></h1>
<div>
  <p align="center">
    Dynamic user segmentation service for AvitoTech internship
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This is repository for `test task` from `AvitoTech`. Service give an `HTTP API` with `JSON` format in `request` and `response`. All project written in Golang language with PostgreSQL database.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This is short instruction about how to start to use this servise in your environment.

### Prerequisites

This is an list of needed software and how to install them.
* Golang language
  ```sh
  sudo apt install golang
  ```
* Docker
  ```sh
  sudo apt install docker docker.io
  ```
  ```sh
  systemctl enable docker
  ```
  ```sh
  systemctl start docker
  ```
* PostgreSQL (via Docker container)
  ```sh
  docker pull postgres
  ```
* Migrate utilite
  ```sh
  curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
  ```
  ```sh
  sudo apt-get install migrate
  ```
  

### Installation

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/github_username/repo_name.git
   ```
3. Install software above
  
4. Write the folow command
   ```
   make build && make run
   ```
5. If service runs for first time, migrations must be done:
   ```
   make migrate
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Whole API Endpoints avalieble at POSTMAN:
```
https://www.postman.com/crimson-resonance-324941/workspace/avitotech-intershiptest/collection/27188643-63e09760-9e57-43ea-8aa0-d9c6520ee754?action=share&creator=27188643
```
As Service installed and runned swagger docs also avalieble at:
```
http://SERVER_HOST:SERVER_PORT/swagger/index.html
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] API Service
- [x] Swagger docs
- [ ] Test covarage
- [x] Extra task one
- [ ] Extra task two
- [x] Extra task three


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Alexey Kirichek - rokirokz@mail.ru

<p align="right">(<a href="#readme-top">back to top</a>)</p>
