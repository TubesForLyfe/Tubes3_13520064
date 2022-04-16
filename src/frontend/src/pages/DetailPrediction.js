import React, { useState } from 'react'
import Axios from 'axios'

const DetailPrediction = () => {
  const [input, setInput] = useState('');
  const [index, setIndex] = useState(1);
  const [predictionList, setPredictionList] = useState([]);

  const getDetailPrediction = (e) => {
    e.preventDefault();
    Axios.post(`${process.env.REACT_APP_DNA_API}/get-detailprediction`, {
      input: input
    }).then((response) => {
      setPredictionList(response.data)
    })
  }

  return (
    <div>
      <input type='text' placeholder='Input date or disease name'
        onChange={(e) => {
          setInput(e)
        }}
      />
      <button onClick={getDetailPrediction}>Process</button>
      <div>
        {predictionList.map((val, key) => {
          return (
            <div>
              <p>{index}. {val.TanggalPrediksi} - {val.NamaPasien} - {val.PenyakitPrediksi} - {val.TingkatKemiripan} - </p>
              {val.Status == 1 && <p>True</p>}
              {val.Status == 0 && <p>False</p>}
              {setIndex(index + 1)}
            </div>
          )
        })}
      </div>
    </div>
  )
}

export default DetailPrediction
