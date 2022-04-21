import { BrowserRouter as Router, Switch, Route, Redirect } from "react-router-dom";

import Landing from "./pages/Landing";
import AddDisease from "./pages/AddDisease";
import DiseasePrediction from "./pages/DiseasePrediction";
import DetailPrediction from "./pages/DetailPrediction";

const RouteManager = () => {
    return (
        <Router>
            <Switch>
                <Route path="/" exact component={Landing} />
                <Route path="/add-disease" component={AddDisease} />
                <Route path="/detail-prediction" component={DetailPrediction} />
                <Route path="/disease-prediction" component={DiseasePrediction} />
            </Switch>
        </Router>
    )
}

export default RouteManager;