# Capstone - Program Catalog
<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#features">Features</a></li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#deploy">Deploy</a>
        <ul>
        <li><a href="#remote">Remote</a></li>
        <li><a href="#local">Local</a></li>
      </ul>
        </li>
        <li><a href="#guide">Guide</a></li>
      </ul>
    </li>
    <ul>
        <li><a href="#erd">ERD Diagram</a></li>
    </ul>
    </li>
    <li><a href="#customization">Customization</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


### Built With

[![Go](https://img.shields.io/badge/golang-007d9c?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

 [![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)](https://react.dev/)


## Usage

This is the University's official Program Catalog, designed to inform students about available academic programs and courses. It includes all necessary requirements to facilitate informed enrollment decisions.


## Key Features
* **Server-Side Rendered Program Catalog (Go Templates):** Delivers a fast and SEO-friendly program catalog by rendering content on the server using Go templates before sending it to the browser.
* **Intuitive Admin Dashboard (React):** Provides a user-friendly administrative interface built with React for faculty to easily manage and update program and course information.
* **Seamless Access Management with SSO:** Integrates Single Sign-On (SSO) for secure and convenient access management for authorized users.
* **Flexible Style Customization:** Offers comprehensive style customization options to align the catalog's appearance with your institution's branding.


## Getting Started
### Deploy Options
#### Remote

- Frontend is hosted at https://capstone.freethegnomes.org/
- Admin dashboard is [https://admincatalog.freethegnomes.org/](https://admin_capstone.freethegnomes.org/)
- API is hosted at https://capstone.freethegnomes.org/api/

---

#### Local

- Install [Docker](https://www.docker.com/products/docker-desktop/)
- [Clone Repo](https://github.com/Tylerchristensen100/CS4900.git)
- In the repo run  `$ docker-compose up`
- Site can be found at http://localhost:3100/
    - Public site is at /
    - Admin dashboard is at /admin/
    - API Docs are at /api/docs


### Guide:
To deploy this application, follow these steps:

1.  **Clone Repository:** Obtain a local copy of the project repository using Git:
    ```bash
    git clone https://github.com/Tylerchristensen100/CS4900
    ```

2.  **Docker Prerequisites:** Ensure that Docker is installed and running on your system. Refer to the official Docker documentation for installation instructions if needed.

3.  **Project Setup:** Open the cloned repository in your preferred code editor.

4.  **Global Styles Configuration:** Customize the global styles for both the public site and the administrative dashboard by modifying the `./core/vars.css` file.

5.  **Environment Variable Configuration:** Configure the necessary environment variables by editing the `.env` file located in the project root.

6.  **Deployment:** Deploy the application using Docker Compose with the following command executed from the project root directory:
    ```bash
    docker compose up -d
    ```
    *(The `-d` flag runs the containers in detached mode.)*



### ERD

![ERD Diagram](https://github.com/Tylerchristensen100/CS4900/blob/main/documentation/ERD.png)



## Customization
* **CSS Styling (`./core/`):** Modify CSS files within this directory to adjust the application's visual design.
* **SSO Configuration (`.env`):** Configure Single Sign-On integration by setting relevant environment variables.
* **Database Configuration (`.env`):** Define database connection parameters via environment variables.
 


## Contributing
We warmly welcome contributions to enhance and improve this project! We are happy to receive pull requests that introduce new features, address bugs, or improve existing functionality.


## Acknowledgments

* [README Template](https://github.com/othneildrew/Best-README-Template)
* [UVU](https://uvu.edu)
* [UVU Web](https://github.com/UVU-WDS)
