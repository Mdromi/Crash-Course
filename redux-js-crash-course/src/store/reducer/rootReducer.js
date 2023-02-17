const initialState = {
  count: 0,
  disable: false
};

const rootReducer = (state = initialState, action) => {
  switch (action.type) {
    case "ADD": {
      return {
        count: state.count + 1,
        disable: false
      };
    }
    case "SUB": {
      if (state.count  > 0) {
        return {
          count: state.count - 1,
          disable: true
        };
      }
    }

    default:
      return state;
  }
};

export default rootReducer;
