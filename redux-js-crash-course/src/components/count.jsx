import React from "react";
import { connect } from "react-redux";

const Count = (props) => {
  return (
    <div>
      <h1>Coming From Reducer</h1>
      <h2>{props.count}</h2>
    </div>
  );
};

const mapStateToProps = (state) => {
  return {
    count: state.count,
    disable: state.disable,
  };
};

export default connect(mapStateToProps)(Count);
