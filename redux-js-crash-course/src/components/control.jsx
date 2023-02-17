import { connect } from "react-redux";

const Control = (props) => {
  console.log("props", props);
  return (
    <div>
      <button onClick={() => props.add()}> + ADD</button>
      <button disabled={props.disable && true} onClick={() => props.sub()}>
        {" "}
        - SUB
      </button>
    </div>
  );
};

const mapDispatchToProps = (dispatch) => {
  return {
    add: () => {
      dispatch({ type: "ADD" });
    },
    sub: () => {
      dispatch({ type: "SUB" });
    },
  };
};

export default connect(null, mapDispatchToProps)(Control);
