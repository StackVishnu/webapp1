
document.addEventListener("DOMContentLoaded", function (){
    const urlParams = new URLSearchParams(window.location.search);
    const productId = urlParams.get('id');
    const role = urlParams.get('role');
    console.log(productId);
    
    fetch(`http://localhost:9090/viewEach/${productId}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            if (data.length > 0) { // Check if data is not empty
                const product = data[0]; // Assuming there's only one product returned
                console.log(product);
                document.getElementById('phone_name').value = product.mName;
                document.getElementById('specs').value = product.mSpecs;
                document.getElementById('price').value = product.mPrice;
            } else {
                console.error('No data found');
            }
        })
        .catch(error => console.error('Error fetching product data:', error));
});

