

addEventListener("DOMContentLoaded", (event) => {
    const cardsSpace = document.querySelectorAll(".cardSpace");
    cardsSpace.forEach(card => {
        card.addEventListener("click", ()=> {
            const spaceID = card.getAttribute("data-space-id");
            console.log("selected space id: ", spaceID);
            if (spaceID == "independent") {
                try {
                    console.log("Independent space selected");
                }catch (err) {
                    console.error("Error selecting independent space: ", err);
                }
            }
            if (spaceID == "my-companies") {
                try {
                    const response = await fetch(`${window.APP_CONFIG.api_url}/`)
                    console.log("My companies space selected");
                }catch (err) {
                    console.error("Error selecting my companies space: ", err);
                }
            }
            if (spaceID == "my-jobs") {
                try {
                    console.log("My jobs space selected");
                }catch (err) {
                    console.error("Error selecting my jobs space: ", err);
                }
            }
                 
        })
    })

})