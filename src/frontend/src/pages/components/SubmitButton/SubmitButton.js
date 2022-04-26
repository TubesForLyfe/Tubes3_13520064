import React, { useState, useEffect } from "react"
import "./SubmitButton.css"
import Axios from 'axios'



const SubmitButton = () => {

    const submitDisease = (event) => {
        var elmt = document.getElementById("message")
        elmt.textContent = "Loading ..."
        var file = document.getElementById('filePenyakit').files[0]
        var namapenyakit = document.getElementById('fieldInputPenyakit').value
        if (file) {
            var reader = new FileReader()
            reader.readAsText(file, "UTF-8")
            reader.onload = function (evt) {
                var content = evt.target.result;
                console.log(namapenyakit)
                console.log(content)

                // post
                event.preventDefault()
                Axios.post(`${process.env.REACT_APP_DNA_API}/submitDisease`, {
                    namaPenyakit: namapenyakit,
                    DNA: content,
                  }).then((response) => {         
                    elmt.textContent = response.data
                    console.log(response.data)
                  })



            }
            reader.onerror = function (evt) {
                // error
                console.log("error")

            }
        }
    }

    return (

        <div class = "SubmitButton colMargin">
            <button id="submitButton" onClick={submitDisease}>Tambahkan Penyakit</button>
        </div>

    )
}


export default SubmitButton;