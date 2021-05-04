import React, {useReducer} from 'react';
import { withCookies } from 'react-cookie';
import axios from 'axios';
import {
    START_FETCH,
    FETCH_SUCCESS,
    ERROR_CATCHED,
    INPUT_EDIT,
    TOGGLE_MODE,
  } from './actionTypes';

const initialState = {
    isLoading: false,
    isLoginView: true,
    error: '',
    credentialsLog: {
      email: '',
      password: '',
    },
    credentialsReg: {
      last_name:  '',
      first_name: '',
      user_name: '',
      email: '',
      password: '',
    },
  };

  type State = typeof initialState

  interface LoginAction {
    type: ActionType;
    payload?: StringKeyObject;
    inputName?: string;
  }
  
  type ActionType = 
    | typeof START_FETCH 
    | typeof FETCH_SUCCESS
    | typeof ERROR_CATCHED
    | typeof INPUT_EDIT
    | typeof TOGGLE_MODE;

  interface LoginInfoInputEvent extends React.FormEvent<HTMLInputElement> {
    target: HTMLInputElement;
  }
  
  interface StringKeyObject {
    [key: string]: string;
  }

  function isValidInputName(arg: any): arg is StringKeyObject{
    return arg !== undefined
  }
  
  const loginReducer = (state: State, action: LoginAction): State => {
    switch (action.type) {
      case START_FETCH: {
        return {
          ...state,
          isLoading: true,
        };
      }
      case FETCH_SUCCESS: {
        return {
          ...state,
          isLoading: false,
        };
      }
      case ERROR_CATCHED: {
        return {
          ...state,
          error: 'Email or Password is not correct!',
          isLoading: false,
        };
      }
      case INPUT_EDIT: {
        // TODO: action.inputNameがオプショナルパラメーター（undifinedが来る可能性）のため、無理やりinputNameのundefinedチェックを入れているが他にやり方がないか検討
        if (isValidInputName(action.inputName)) {
          return {
            ...state,
            [action.inputName]: action.payload,
            error: '',
          };
        } else {
          return state
        }
      }
      case TOGGLE_MODE: {
        return {
          ...state,
          isLoginView: !state.isLoginView,
        };
      }
      default:
        return state;
    }
  };

const Login: React.FC = (props: any) => {
    const [state, dispatch] = useReducer(loginReducer, initialState);

    const inputChangedLog = () => (event: LoginInfoInputEvent) => {
      const cred: StringKeyObject = state.credentialsLog;
      cred[event.target.name] = event.target.value;
      dispatch({
        type: INPUT_EDIT,
        inputName: 'state.credentialLog',
        payload: cred,
      });
    };
  
    const inputChangedReg = () => (event: LoginInfoInputEvent) => {
      const cred: StringKeyObject = state.credentialsReg;
      cred[event.target.name] = event.target.value;
      dispatch({
        type: INPUT_EDIT,
        inputName: 'state.credentialReg',
        payload: cred,
      });
    };

    const login = async(event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (state.isLoginView) {
          try {
            dispatch({ type: START_FETCH });
            const res = await axios.post(
              'http://127.0.0.1:9090/api/v1/login',
              state.credentialsLog,
              {
                headers: { 'content-type': 'application/json' },
              }
            );
            props.cookies.set('current-token', res.data.response.token);
            res.data.response.token
              ? (window.location.href = '/profiles')
              : (window.location.href = '/');
            dispatch({ type: FETCH_SUCCESS });
          } catch {
            dispatch({ type: ERROR_CATCHED });
          }
        } else {
          try {
            dispatch({ type: START_FETCH });
            await axios.post(
              'http://127.0.0.1:9090/api/v1/signup',
              state.credentialsReg,
              {
                headers: { 'content-type': 'application/json' },
              }
            );
            dispatch({ type: FETCH_SUCCESS });
            dispatch({ type: TOGGLE_MODE });
          } catch {
            dispatch({ type: ERROR_CATCHED });
          }
        }
    };

    const toggleView = () => {
      dispatch({ type: TOGGLE_MODE });
    };

    return (
        <form onSubmit={login}>
            {state.isLoginView ? 'Login' : 'Register'}
            {state.isLoginView ? (
                <div>
                    <label>
                        Email:
                        <input type="text" name="email" value={state.credentialsLog.email} onChange={inputChangedLog()}/>
                        <br/>
                        Password:
                        <input type="text" name="password" value={state.credentialsLog.password} onChange={inputChangedLog()}/>
                        <br/>
                    </label>
                    <input type="submit" value="Submit"/>
                </div>
            ) : (
                <div>
                    <label>
                        LastName:
                        <input type="text" name="last_name" value={state.credentialsReg.last_name} onChange={inputChangedReg()}/>
                        <br/>
                        FirstName:
                        <input type="text" name="first_name" value={state.credentialsReg.first_name} onChange={inputChangedReg()}/>
                        <br/>
                        UserName:
                        <input type="text" name="user_name" value={state.credentialsReg.user_name} onChange={inputChangedReg()}/>
                        <br/>
                        Email:
                        <input type="text" name="email" value={state.credentialsReg.email} onChange={inputChangedReg()}/>
                        <br/>
                        Password:
                        <input type="text" name="password" value={state.credentialsReg.password} onChange={inputChangedReg()}/>
                        <br/>
                    </label>
                    <input type="submit" value="Submit"/></div>
                )
            }
            <span>{state.error}</span>


            <span onClick={() => toggleView()}>
                {state.isLoginView ? 'Create Account ?' : 'Back to Login ?'}
            </span>
        </form>
    );
};

export default withCookies(Login);