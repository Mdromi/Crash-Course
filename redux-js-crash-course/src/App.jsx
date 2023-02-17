import { Provider } from "react-redux";
import { createStore } from "redux";
import "./App.css";

import Control from "./components/control";
import Count from "./components/count";
import store from "./store";

function App() {
  return (
    <Provider store={store}>
      <div className="App">
        <Count />
        <Control />
      </div>
    </Provider>
  );
}

export default App;

// Store -> Will hold all of our application data/state
// Reducer -> A function  who return a specific amount of state of data
// Actions -> Event occurs => {type: 'Something' payload: ''}
// Dispatch
// Subscriber

// 1. create a reducer
// 2. create store with the help reducer
// 3. Now we can subscribe
// 3. dispatch[action]

// React will be responsible for only view layer
// Redux will be responsible fot only data layer

function Redux_Concept() {
  // A reducer function must have two parameters
  // State, Action
  const reducer = (state = {}, action) => {
    if (action.type === "A") return { ...state, A: "I am A" };
    if (action.type === "B") return { ...state, B: "I am B" };
    if (action.type === "C") return { ...state, C: "I am C" };
  };

  const store = createStore(reducer);

  store.subscribe(() => {
    console.log(store.getState().A);
  });

  store.dispatch({ type: "A" });
  store.dispatch({ type: "B" });
  store.dispatch({ type: "Something" });
  store.dispatch({ type: "Something" });

  return <div className="App"></div>;
}
