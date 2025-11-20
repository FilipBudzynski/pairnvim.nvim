local M = {}
require("host")
Addr = "127.0.0.1:6666"
local function handle_callback(_, data)
    vim.notify(table.concat(data, "\n"))
end

function M.connect()
    local chan = vim.fn.sockconnect("tcp", Addr, { on_data = handle_callback, data_buffered = false })
    if chan == 0 then
        vim.notify("did not connect to a channel", 1)
        print("did not connect to a channel")
        return
    end
    vim.notify("successfuly connected to channel: " .. chan, 1)
    vim.fn.chansend(chan, "witam\n")
end

vim.api.nvim_create_user_command("Pair", M.connect, {})

return M
