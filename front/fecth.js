const estado = document.getElementById("backendStatus");
fetch("http://localhost:8080/ping")
    .then(res => res.json())
    .then(data => {
        estado.textContent = data.message;
    })
    .catch(error => {
        estado.textContent = "Error: " + error.message;
        console.error("Fetch error:", error);
    });