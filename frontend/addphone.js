function insert(event) {
    event.preventDefault();

    // Obtain all form data
    const phoneName = document.getElementById('phone_name').value;
    const specs = document.getElementById('specs').value;
    const price = document.getElementById('price').value;
    const imageFile = document.getElementById('image').files[0];

    // Create FormData object
    const formData = new FormData();
    formData.append('phone_name', phoneName);
    formData.append('specs', specs);
    formData.append('price', price);
    formData.append('image', imageFile);

    // Send FormData to API endpoint
    fetch('http://localhost:9090/insert', {
        method: 'POST',
        body: formData
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        // Handle successful response
        alert("Product added successfully");
        window.location.href ="admin.html"
    })
    .catch(error => {
        // Handle errors
        console.error("Error editing product:", error);
            // Optionally, display an error message to the user
            console.log("Error name:", error.name);
            console.log("Error message:", error.message);
            console.log("Error stack trace:", error.stack);
            alert("Error editing product. Please try again later.");
    });
}
