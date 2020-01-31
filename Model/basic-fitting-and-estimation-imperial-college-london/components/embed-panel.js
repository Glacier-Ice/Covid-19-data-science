const React = require('react');
const utils = require('./utils');

class CustomComponent extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      revealed: false,
      enabled: false,
      rerender: 0,
      isMobile: false
    };

    let lastState = this.state.enabled;
    setInterval(() => {
      if (this.state.enabled !== lastState) {
        // setTimeout(() => {
          // this.setState({
          //   rerender: this.state.rerender + 1
          // })
        // }, 2000)
      }
      lastState = this.state.enabled;
    }, 1000)
  }

  componentDidMount() {
    this.setState({
      isMobile: utils.isMobile()
    })
  }

  handleClick() {
    this.setState({ enabled: !this.state.enabled });
  }

  render() {
    const { idyll, hasError, updateProps, children, ...props } = this.props;
    const { width = 1, height = 1, row = 1, column = 1, borderColor, color = '#fff', svg = 'small-dog', url = 'https://jheer.github.io/barnes-hut/' } = props;
    const styles = {
      gridColumn: this.state.isMobile ? `1` : `${column} / span ${width}`,
      gridRow: this.state.isMobile ? `span 1` : `${row} / span ${height}`,
      backgroundColor: color,
      borderColor: borderColor ? borderColor : undefined
    }
    const svgStyles = {
      // opacity: this.state.revealed ? 1 : 0.25,
      // filter: this.state.revealed ? `blur(0)` : `blur(0.25rem)`
    }
    const wrapperStyles = this.state.enabled ? {
      position: 'fixed',
      left: 0,
      right: 0,
      top: 0,
      bottom: 0,
      background: 'white',
      height: '100%',
      width: '100%',
      zIndex: 999,
      // overflow: 'auto',
      transition: `all 0.5s`
    } : {
      position: 'relative',
      height: '100%',
      width: '100%',
      transition: `all 0.5s`
    }
    return (
      <div className={'panel embed'} style={styles} onClick={this.handleClick.bind(this)}>
        <div className={`fullscreen ${this.state.enabled ? ' fullscreen-enabled' : ''}`} style={wrapperStyles}>
          <iframe key={this.state.rerender} src={url} style={{display: 'block', position: 'absolute', top: 0, left: 0, bottom: 0, right: 0}} />
          {
            this.state.enabled ?
            <div style={{position: 'absolute', right: 20, top: 20, padding: 10, background: `rgba(0, 0, 0, 0.8)`, color: 'white'}}>Exit</div>
            :
            <div style={{position: 'absolute', right: 0, top: 0, left: 0, bottom: 0, background: `rgba(0, 0, 0, 0)`}} />
          }
        </div>
      </div>
      // <div className={this.state.revealed ? 'panel revealed' : 'panel'} style={styles} onClick={() => this.setState({ revealed: true })}>
      //   <SVG
      //       src={`static/svg/${svg}.svg`}
      //   />
      // </div>
    );
  }
}

module.exports = CustomComponent;
