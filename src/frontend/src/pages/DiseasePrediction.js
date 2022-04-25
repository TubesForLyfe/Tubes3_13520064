import React, { useState } from 'react'
import './DiseasePrediction.css'
import Axios from 'axios'



const DiseasePrediction = () => {
  const [inputName, setInputName] = useState('');
  const [inputFile, setInputFile] = useState();
  const [inputPenyakit, setInputPenyakit] = useState('');

  const [hasilTes, setHasilTes] = useState([]);

  const getDiseasePrediction = (e) => {
    e.preventDefault();

    const formData = new FormData();
    formData.append('file', inputFile);
    formData.append('inputName', inputName);
    formData.append('inputPenyakit', inputPenyakit);
    formData.append('tanggal', new Date().toLocaleString().split(" ")[0]);

    Axios.post(`${process.env.REACT_APP_DNA_API}/get-diseaseprediction`, formData)
    .then((response) => {
      setHasilTes(response.data)
    })
  }

  return (
    <div>
      <div id="tesDNA">
        <div>
          <h3> Tes DNA </h3>
          <div className = "column side">
            <p>Nama Penguna:</p>
            <input className="inputBox" type='text' placeholder='Nama pengguna...' required 
            onChange={(e) => {
              setInputName(e.target.value)
            }}></input>
          </div>
          <div className = "column middle">
            <p>Sequence DNA:</p>
            <input className="inputBox" type="file" accept=".txt" required
            onChange={(e) => {
              setInputFile(e.target.files[0])
            }}></input>
          </div>
          <div className = "column side">
            <p>Prediksi Penyakit:</p>
            <input className="inputBox" type='text' placeholder='Nama Penyakit...' required
            onChange={(e) => {
              setInputPenyakit(e.target.value)
            }}></input>
          </div>
          <div>
            <button onClick={getDiseasePrediction}>Submit</button>
          </div>
        </div>
        <div id="hasilTes">
          <h3> Hasil </h3>
          {hasilTes.map((val, key) => {
            return (
              <div>
              {(val.Status === 0 || val.Status === 1)&&
                 <div>
                <p className="Same"> {val.TanggalPrediksi} -</p>
                <p className="Same">&nbsp; {val.NamaPasien} -</p>
                <p className="Same">&nbsp; {val.TingkatKemiripan}% -</p>
                <p className="Same">&nbsp; {val.PenyakitPrediksi} -</p>
                {val.Status == 1 && <p className="Same">&nbsp; True</p>}
                {val.Status == 0 && <p className="Same">&nbsp; False</p>}
                </div>
              } 
              {val.Status === -1 &&
                <p className="Same">DNA Tidak Valid!</p>
              }
              {val.Status === -2 &&
                <p className="Same">Penyakit Tidak Ditemukan</p>
              }
              {val.Status === -3 &&
                <p className="Same">Penyakit Tidak Ditemukan</p>
              }
              </div>
            )
          })}
        </div>
      </div>
      
    </div>
  )
}

export default DiseasePrediction
