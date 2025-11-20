local M = {}

M.buffers = {}
M.server_job = nil
M.channel = ""

function M.send(self, data)
    vim.fn.chansend(self.channel, data)
end

local function scrap_buffer()
    local lines = vim.api.nvim_buf_get_lines(0, 0, -1, true)
    local text = table.concat(lines, "\n")
    M.buffers.main = text
end

local function handle_callback(_, data)
    vim.notify(table.concat(data, "\n"))
end

local function connect_to_server()
    if M.server_job.pid == nil then
        vim.notify("Server is not running", vim.log.levels.ERROR)
        return
    end

    M.channel = vim.fn.sockconnect("tcp", "127.0.0.1:6666", {
        rpc = false,
        on_data = handle_callback,
        data_buffered = false,
    })

    if M.channel == 0 then
        print("Host failed to connect")
    else
        print("Host connected to TCP socket:", M.channel)
    end

    vim.fn.chansend(M.channel, "Witam z serwera")
end

vim.api.nvim_create_user_command("Start", function()
    -- run golang server
    M.server_job = vim.system({ "go", "run", "server/main.go" }, { detach = true })
    vim.notify(tostring(M.server_job.pid))

    -- connect to it
    local timer = vim.uv.new_timer()
    timer:start(2000, 0, vim.schedule_wrap(connect_to_server))

    --M:send("witam z serwera")
end, {})
return M
