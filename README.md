<html>
<body>
    <h1>JWT Backend</h1>
    <p><strong>Simple JWT User Authentication</strong></p>
    
  <h2>Overview</h2>
    <p>
        This is a simple backend application implementing user authentication using JSON Web Tokens (JWT). 
        Built with Golang, it leverages <a href="https://gin-gonic.com/">Gin</a> as the web framework, 
        <a href="https://gorm.io/">Gorm</a> as the ORM, and PostgreSQL as the primary database. 
        SQLite is used for testing purposes, and the application can be easily deployed using Docker.
    </p>

  <h2>Features</h2>
    <ul>
        <li>User registration</li>
        <li>User login</li>
        <li>JWT-based authentication</li>
        <li>Integration tests</li>
        <li>Dockerized deployment</li>
    </ul>
  <h2>Technologies Used</h2>
    <ul>
        <li><strong>Programming Language:</strong> Golang</li>
        <li><strong>Web Framework:</strong> Gin</li>
        <li><strong>ORM:</strong> Gorm</li>
        <li><strong>Primary Database:</strong> PostgreSQL</li>
        <li><strong>Database for Testing:</strong> SQLite</li>
        <li><strong>Testing:</strong> Integration tests included</li>
        <li><strong>Containerization:</strong> Docker</li>
    </ul>

  <h2>Getting Started</h2>
    <h3>Prerequisites</h3>
    <ul>
        <li>Go 1.20 or higher installed</li>
        <li>Git installed</li>
        <li>Docker installed</li>
    </ul>

  <h3>Installation</h3>
    <pre>
<code># Clone the repository
git clone https://github.com/Dnreikronos/jwt-backend.git

# Navigate to the project directory
cd jwt-backend

# Install dependencies
go mod tidy

# Build and run with Docker
docker-compose up --build

# Run tests
make run

# Start the server without Docker
make run
    </code>
    </pre>

  <h2>Usage</h2>
    <h3>Routes</h3>
    <ul>
        <li><code>POST /register</code> - User registration</li>
        <li><code>POST /login</code> - User login and JWT token generation</li>
        <li><code>GET /protected</code> - Access protected route (requires JWT)</li>
    </ul>

   <h2>Testing</h2>
    <p>
        The application includes integration tests that use SQLite for testing purposes. 
        Run the tests using the following command:
    </p>
    <pre>
<code>make tests
    </code>
    </pre>

  <h2>Contributing</h2>
    <p>Contributions are welcome! Feel free to fork the repository and submit a pull request.</p>

   <h2>License</h2>
    <p>This project is licensed under the <strong>MIT License</strong>.</p>

  <footer>
        <p>Developed by <a href="https://github.com/Dnreikronos">Dnreikronos</a></p>
    </footer>
</body>
</html>
