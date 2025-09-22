document.getElementById('registerForm').addEventListener('submit', async function (e) {
  const formContainer = document.getElementById('formContainer')
  const successfulContainer = document.getElementById('successfulContainer')
  const errorContainer = document.getElementById('errorContainer')

  e.preventDefault()

  const params = new URLSearchParams(window.location.search)
  const token = params.get('token')

  const email = document.getElementById('email').value
  const succesfulMessage = document.getElementById('succesfulMessage')
  const errorMessage = document.getElementById('errorMessage')

  formContainer.style.display = 'none'
  try {
    const response = await fetch(window.APP_CONFIG.url_register, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ email, token })

    })
    if (!response.ok) throw new Error('Ocurrio algun problema')

    const data = await response.json()
    successfulContainer.style.display = 'block'
  } catch {
    errorContainer.style.display = 'block'
  }
})
