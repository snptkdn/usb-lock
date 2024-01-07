<script>
  import logo from "./assets/images/logo_transparent.png";
  import { GetUsbDevices } from "../wailsjs/go/main/App.js";

  let name;
  let records = [];

  function getUSB() {
    GetUsbDevices().then((result) => records = result);
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src={logo} />
  <div class="input-box" id="input">
    <button class="btn" on:click={getUSB}>USB機器を探索する</button>
  </div>
  <div class="usb-table-area">
    <table class="table">
      <thead>
        <th class="td">モデル</th>
        <th class="td">ディスクID</th>
        <th class="td">容量</th>
        <th class="td-button">状態切り替え</th>
      </thead>
      {#each records as record}
        <tbody>
          <tr>
            <td class="td">{record.Model}</td>
            <td class="td">{record.Index}</td>
            <td class="td">{record.Size}</td>
            <td class="td-button">
              <button class="btn-lock">{record.button}</button>
            </td>
          </tr>
        </tbody>
      {/each}
    </table>
  </div>
</main>

<style>
  #logo {
    display: block;
    width: 40%;
    height: 40%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .usb-table-area {
    margin: 30px 0 0;
    justify-content: center;
    display: flex;
  }

  .table {
    border: 1px solid white;
    border-collapse: collapse;
  }

  .table .td {
    border: 1px solid white;
    width: 100px;
  }

  .td-button {
    border: 1px solid white;
    width: 100px;
  }

  .btn {
    width: auto;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .btn-lock {
    width: 100px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }
</style>
