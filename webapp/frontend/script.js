document.addEventListener("DOMContentLoaded",function (){
    if (window.location.pathname == "/home/vishnurathan/Desktop/webapp/frontend/user.html") {
        fetch("http://localhost:9090/viewAll")
    .then(response => {
        if (!response.ok){
            throw new Error('Network response was not ok');
        }
       return response.json()
    })
    .then(data => {
        
        data.forEach(element => {
            createCard(element);
        });
    })
    .catch(error => console.error('Error fetching product data:', error))
    }
    
})

function createCard(el){
    const container = document.getElementById("pCard")
    if (container.children.length === 0){
        const card = document.createElement("div")
        card.classList.add("card","mb-3")
        card.innerHTML = `
                    <div class="card-body">
                    <div class="card-body d-flex justify-content-center align-items-center">
                      <img src="./assets/phone.jpeg" class="card-img-top" alt="..." style="width: 100px; height: 200px;">
                    </div>

                    <div class="card-body d-flex justify-content-center align-items-center">
                      <h5 class="card-title">${el.mid}</h5>                      
                    </div>
                    
                    <div class="card-body d-flex justify-content-center align-items-center">
                    <a href="/home/vishnurathan/Desktop/webapp/frontend/detailedview.html?id=${el.mid}" class="btn btn-primary view-btn" >VIEW</a> 
                    </div>

                    </div>
        `
        container.appendChild(card)
    }
    
}


