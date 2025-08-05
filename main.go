package main

import (
	"runtime"

	glfw "github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	defer window.Destroy()

	procAddr := glfw.GetVulkanGetInstanceProcAddress()
	if procAddr == nil {
		panic("GetInstanceProcAddress is nil")
	}
	vk.SetGetInstanceProcAddr(procAddr)

	if err := vk.Init(); err != nil {
		panic(err)
	}

	vkInstance := createInstance()
	defer vk.DestroyInstance(vkInstance, nil)

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func createInstance() vk.Instance {
	appInfo := vk.ApplicationInfo{
		SType:              vk.StructureTypeApplicationInfo,
		PApplicationName:   "MIG",
		ApplicationVersion: vk.MakeVersion(0, 1, 0),
		PEngineName:        "No Engine",
		EngineVersion:      vk.MakeVersion(1, 0, 0),
		ApiVersion:         vk.ApiVersion10,
	}

	createInfo := vk.InstanceCreateInfo{
		SType:                   vk.StructureTypeInstanceCreateInfo,
		PApplicationInfo:        &appInfo,
		EnabledExtensionCount:   0,
		PpEnabledExtensionNames: glfw.GetCurrentContext().GetRequiredInstanceExtensions(),
		EnabledLayerCount:       0,
	}

	var vkInstance vk.Instance

	result := vk.CreateInstance(&createInfo, nil, &vkInstance)
	if result != vk.Success {
		panic("failed to create vulkan instance")
	}

	return vkInstance
}
