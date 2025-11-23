local M = {}

M.buffers = {}
M.server_job = nil
M.channel = ""
M.server_dir = vim.fs.normalize(debug.getinfo(1, "S").source:sub(2):match("(.*/)") .. "../server")

local function listen_to_server(_, data)
	if data then
		print("Server: " .. data)
	end
end

--- @class opts
--- @field logger fun(_, data: string|string[])
M.opts = {
	logger = listen_to_server,
}

--- @param opts opts
function M.setup(opts)
	for k, v in ipairs(opts) do
		M.opts[k] = v
	end
end

local SERVER_FILENAME = "server"

function M.send(self, data)
	vim.fn.chansend(self.channel, data)
end

local function handle_callback(_, data)
	vim.notify(table.concat(data, "\n"))
end

local function cleanup()
	if M.server_job and M.server_job.pid then
		M.server_job:kill(9)
		M.server_job = nil
	end

	if M.channel and M.channel ~= 0 then
		vim.fn.chanclose(M.channel)
	end
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
		vim.fn.chansend(M.channel, "Witam z serwera")
	end
end

function M.run_server()
	if M.server_job ~= nil then
		return
	end

	vim.system({ "go", "build", "-o", SERVER_FILENAME }, {
		cwd = M.server_dir,
	}):wait()

	M.server_job = vim.system({ "./" .. SERVER_FILENAME }, {
		cwd = M.server_dir,
		stdout = listen_to_server,
		stderr = listen_to_server,
		detach = true,
	})
end

vim.api.nvim_create_user_command("Start", function()
	M.run_server()
	vim.notify("server running" .. tostring(M.server_job.pid))

	if M.timer then
		M.timer:stop()
	end
	M.timer = vim.uv.new_timer()
	M.timer:start(500, 0, vim.schedule_wrap(connect_to_server))
end, {})

local server_cleanup_group = vim.api.nvim_create_augroup("ServerCleanup", { clear = true })
vim.api.nvim_create_autocmd("VimLeavePre", {
	group = server_cleanup_group,
	callback = cleanup,
	desc = "Cleanup server on exit",
})

return M
