import React from 'react'
import { Link } from 'react-router-dom'

const Landing = () => {
  return (
    <div>
      <p>Welcome to DNA Application</p>
      <p><Link to="/add-disease">Add Disease</Link></p>
      <p><Link to="/disease-prediction">Disease Prediction</Link></p>
      <p><Link to="/detail-prediction">Detail Prediction</Link></p>
    </div>
  )
}

export default Landing
