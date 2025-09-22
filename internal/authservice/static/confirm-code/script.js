const params = new URLSearchParams(window.location.search)
const token = params.get('token')

document.addEventListener('DOMContentLoaded', () => {
  const confirmForm = document.getElementById('confirmForm')
  const tryAgain = document.getElementById('tryAgain')
  const password1 = document.getElementById('password1')
  const password2 = document.getElementById('password2')

  const formContainer = document.getElementById('formContainer')
  const errorContainer = document.querySelector('.errorContainer')
  const successContainer = document.querySelector('.successContainer')

  function validarContrasenia (pass) {
    return [
      pass.length >= 8, // Minimo 8 letras
      /[A-Z]/.test(pass), // Al menos una letrea mayuscula
      /[a-z]/.test(pass), // Al menos una letra minuscula
      /[\d]/.test(pass), // Al menos un numero
      /[\W_]/.test(pass)
    ]
  }

  confirmForm.addEventListener('submit', async function (e) {
    e.preventDefault()

    const errorMessage = document.getElementById('errorMessage')
    const successMessage = document.getElementById('successMessage')

    const pass1 = password1.value.trim()
    const pass2 = password2.value.trim()
    const validations = validarContrasenia(pass1)

    updateIcon('check-length', validations[0])
    updateIcon('check-uppercase', validations[1])
    updateIcon('check-lowercase', validations[2])
    updateIcon('check-number', validations[3])
    updateIcon('check-special', validations[4])

    const requirementsOk = validations.every(Boolean)

    if (!requirementsOk) {
      return alert('La contrasenia no cumple todos los requisitos.')
    }

    if (pass1 !== pass2) {
      return alert('Las contrasenias no coinciden.')
    }

    formContainer.style.display = 'none'

    try {
      const response = await fetch(window.APP_CONFIG.url_confirm, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ contrasenia: pass1, token })
      })

      if (!response.ok) throw new Error('Error al crear el usuario')

      successContainer.style.display = 'block'
      successMessage.textContent = 'Registro completado!'
    } catch (error) {
      console.log(error)
      errorContainer.style.display = 'block'
      errorMessage.textContent = 'Error en el servidor!'
    }
  })

  tryAgain.addEventListener('click', () => {
    formContainer.style.display = 'block'
    errorContainer.style.display = 'none'
  })
})

function updateIcon (id, condition) {
  const icon = document.getElementById(id)
  icon.classList.remove('fa-circle-check', 'fa-circle-xmark')

  if (condition) {
    icon.classList.add('fa-circle-check')
    icon.style.color = 'lightgreen'
  } else {
    icon.classList.add('fa-circle-xmark')
    icon.style.color = 'lightcoral'
  }
}
