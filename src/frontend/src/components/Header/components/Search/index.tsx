import './index.scss'
// type SearchProps = {
// };
function Search(/* props: SearchProps */) {
  // const {} = props;
  return (
    <div className='c-header-search'>
      <input
        type='text'
        placeholder='Search'
        className='c-header-search-input'
      />
    </div>
  )
}
export default Search
